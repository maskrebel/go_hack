with root as (
  select
    N
  from
    BST
  where
    P is null
),
inner1 as (
  select
    N
  from
    BST
  where
    P in (
      select
        N
      from
        root
    )
)
select
  N,
  case when P is Null then 'Root' when P in (
    select
      N
    from
      root
  ) then 'Inner' when P in (
    select
      N
    from
      inner1
  ) then 'Inner' else 'Leaf' end
from
  BST
order by
  N;
