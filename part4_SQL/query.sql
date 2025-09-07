-- 1. Basic query
GO

SELECT *
FROM books
WHERE price > 30

-- 2. Aggregation
GO 

SELECT user_id, SUM(amount) as TotalAmountSpent
FROM orders
GROUP BY user_id