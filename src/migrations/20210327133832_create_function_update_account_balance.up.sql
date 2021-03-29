create or replace function update_account_balance(accounts_id_update int, source_type_update text) returns setof transactions as
$$
DECLARE
    --get transactions rows
    tr               transactions%rowtype;
    --user balance
    user_balance     decimal;
    --sum for blocked
    sum_blocked      decimal;
    --the amount to be reset
    -- resolve - operation ready success
    -- blocked - operation blocked success
    -- adjustment - operation adjustment
    status_update    text;
    -- If the glare balance exceeds the user's balance, then we catch the error
    catch_exceptions bool;
begin
    catch_exceptions = false;
    --get user balance
    user_balance = balance FROM accounts WHERE accounts.account_id = accounts_id_update;
    -- calc sum for revert
    sum_blocked = sum(amount)
                  FROM (SELECT *
                        FROM transactions
                        WHERE ((transaction_id % 2) <> 0)
                        ORDER BY transaction_id DESC
                        LIMIT 10) sub;


    IF sum_blocked > user_balance THEN
        catch_exceptions = true;
        status_update = 'adjustment';
    END IF;

    FOR tr IN select *
              FROM (SELECT *
                    FROM transactions
                    WHERE ((transaction_id % 2) <> 0)
                    ORDER BY transaction_id DESC
                    LIMIT 10) sub
              WHERE status = 'pending'
              ORDER BY transaction_id
        LOOP
            UPDATE transactions
            SET status      = 'blocked',
                source_type = source_type_update
            WHERE account_id = accounts_id_update
              AND transaction_id = tr.transaction_id;


            UPDATE accounts
            SET balance = balance + (tr.amount * -1)
            WHERE account_id = accounts_id_update;

            RETURN NEXT tr;
        END LOOP;

    FOR tr IN
        SELECT *
        FROM transactions
        WHERE status = 'pending'
        ORDER BY transaction_id DESC
        LIMIT 10
        LOOP
            UPDATE transactions
            SET status      = 'resolve',
                source_type = source_type_update
            WHERE account_id = accounts_id_update
              AND transaction_id = tr.transaction_id;
        END LOOP;


    IF catch_exceptions THEN
        RAISE EXCEPTION 'Invalid balance for update.'
            USING HINT = 'The account does not have enough money',
                ERRCODE = 12882;
    END IF;

    RETURN;
end;

$$
    language plpgsql;