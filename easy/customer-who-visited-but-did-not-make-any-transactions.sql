-- 1581. 进店却未进行过交易的顾客
-- https://leetcode.cn/problems/customer-who-visited-but-did-not-make-any-transactions/description/

/*表：Visits

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| visit_id    | int     |
| customer_id | int     |
+-------------+---------+
visit_id 是该表中具有唯一值的列。
该表包含有关光临过购物中心的顾客的信息。*/

/*表：Transactions

+----------------+---------+
| Column Name    | Type    |
+----------------+---------+
| transaction_id | int     |
| visit_id       | int     |
| amount         | int     |
+----------------+---------+
transaction_id 是该表中具有唯一值的列。
此表包含 visit_id 期间进行的交易的信息*/

SELECT
    customer_id,
    count( customer_id ) AS count_no_trans
FROM
    Visits
WHERE
    visit_id NOT IN ( SELECT DISTINCT ( visit_id ) AS visit_ids FROM Transactions )
GROUP BY
    customer_id

