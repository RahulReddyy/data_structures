EmployeeID	FirstName	LastName	Department	Salary	HireDate
1	John	Doe			Sales	45000	2022-03-15
2	Jane	Smith		HR	55000	2021-12-10
3	Mark	Johnson		IT	60000	2022-06-20
4	Lisa	Williams	 Sales	48000	2023-01-05
5	David	Brown		IT	52000	2022-09-30


select max(salary) as highsalary from employess 


count of employees from each Department


select count(sales) as sdep, count(it) as itdep,

select * from employees where FirstName like = "%d"