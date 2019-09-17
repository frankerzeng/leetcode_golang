/*
1179. Reformat Department Table
Table: Department

+---------------+---------+
| Column Name   | Type    |
+---------------+---------+
| id            | int     |
| revenue       | int     |
| month         | varchar |
+---------------+---------+
(id, month) is the primary key of this table.
The table has information about the revenue of each department per month.
The month has values in ["Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"].


Write an SQL query to reformat the table such that there is a department id column and a revenue column for each month.

The query result format is in the following example:

Department table:
+------+---------+-------+
| id   | revenue | month |
+------+---------+-------+
| 1    | 8000    | Jan   |
| 2    | 9000    | Jan   |
| 3    | 10000   | Feb   |
| 1    | 7000    | Feb   |
| 1    | 6000    | Mar   |
+------+---------+-------+

Result table:
+------+-------------+-------------+-------------+-----+-------------+
| id   | Jan_Revenue | Feb_Revenue | Mar_Revenue | ... | Dec_Revenue |
+------+-------------+-------------+-------------+-----+-------------+
| 1    | 8000        | 7000        | 6000        | ... | null        |
| 2    | 9000        | null        | null        | ... | null        |
| 3    | null        | 10000       | null        | ... | null        |
+------+-------------+-------------+-------------+-----+-------------+

Note that the result table has 13 columns (1 for the department id + 12 for the months).
*/
# Write your MySQL query statement below
SELECT
    id,
    SUM( IF( MONTH = 'Jan', revenue, NULL ) ) AS Jan_Revenue,
    SUM( IF( MONTH = 'Feb', revenue, NULL ) ) AS Feb_Revenue,
    SUM( IF( MONTH = 'Mar', revenue, NULL ) ) AS Mar_Revenue,
    SUM( IF( MONTH = 'Apr', revenue, NULL ) ) AS Apr_Revenue,
    SUM( IF( MONTH = 'May', revenue, NULL ) ) AS May_Revenue,
    SUM( IF( MONTH = 'Jun', revenue, NULL ) ) AS Jun_Revenue,
    SUM( IF( MONTH = 'Jul', revenue, NULL ) ) AS Jul_Revenue,
    SUM( IF( MONTH = 'Aug', revenue, NULL ) ) AS Aug_Revenue,
    SUM( IF( MONTH = 'Sep', revenue, NULL ) ) AS Sep_Revenue,
    SUM( IF( MONTH = 'Oct', revenue, NULL ) ) AS Oct_Revenue,
    SUM( IF( MONTH = 'Nov', revenue, NULL ) ) AS Nov_Revenue,
    SUM( IF( MONTH = 'Dec', revenue, NULL ) ) AS Dec_Revenue
FROM
    Department
GROUP BY
    id;