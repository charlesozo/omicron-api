-- name: CreateUserSubscription :one
INSERT INTO Subscription(subscription_id, userid, whatsapp_number, expiry_date, registered_user)
VALUES($1, $2, $3, $4, $5)
RETURNING *;
--
-- name: InputPlan :exec
INSERT INTO Plan (plan_name, description, price, duration_days) VALUES
('Monthly Plan', 'Unlock premium access for one month.', 6.99, 30),
('Annual Plan', 'Unlock premium access for one year at a discounted rate.', 71.99, 365);
--
-- name: PaymentSetup :one
INSERT INTO Payment (payment_id, userid, amount, plan_id, payment_date, expiry_date, payment_method)
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING *;