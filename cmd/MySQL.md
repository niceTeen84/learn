# MySQL 8

## 生成一列数据

- 8 以前的写法

```sql
select * from 
(
   select 1 as val
   union all
   select 2
   union all 
   select 3 
)b;
```

- 8 之后

```sql
select * from (
	values 
	row(1), 
	row(2), 
	row(3)
)b;
```