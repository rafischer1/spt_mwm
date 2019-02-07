BEGIN;

  CREATE TABLE category
  (
    id SERIAL PRIMARY KEY,
    parent_id INTEGER REFERENCES category(id)
    DEFERRABLE,
    name TEXT NOT NULL UNIQUE );

    SET CONSTRAINTS
    ALL DEFERRED;

  INSERT INTO category
  VALUES
    (1, NULL, 'animal');
  INSERT INTO category
  VALUES
    (2, NULL, 'mineral');
  INSERT INTO category
  VALUES
    (3, NULL, 'vegetable');
  INSERT INTO category
  VALUES
    (4, 1, 'dog');
  INSERT INTO category
  VALUES
    (5, 1, 'cat');
  INSERT INTO category
  VALUES
    (6, 4, 'doberman');
  INSERT INTO category
  VALUES
    (7, 4, 'dachshund');
  INSERT INTO category
  VALUES
    (8, 3, 'carrot');
  INSERT INTO category
  VALUES
    (9, 3, 'lettuce');
  INSERT INTO category
  VALUES
    (10, 11, 'paradox1');
  INSERT INTO category
  VALUES
    (11, 10, 'paradox2');
    
  SELECT setval('category_id_seq', (select max(id)
    from category));

-- WITH RECURSIVE last_run (parent_id, id_list,name_list) AS (
--         SELECT string_agg(parent_id::text, ', ') AS parent_id, string_agg(id::text, ', ') AS id_list, string_agg(name, ', ') AS name_list FROM category
--         UNION ALL
--         SELECT parent_id, id_list, name_list FROM last_run
--         WHERE parent_id = id_list
--         -- ORDER BY id_list
-- )

-- WITH RECURSIVE last_run (parent_id, id_list, name_list) AS (
--   SELECT parent_id, ARRAY[id], ARRAY[name] AS name_list from category
--   -- WHERE parent_id = category.id
--   UNION ALL
--   SELECT category.parent_id, id_list || category.id, name_list || ', ' || category.name
--   FROM last_run, category
--   WHERE category.parent_id = last_run.parent_id 
-- )

-- SELECT id_list, array_to_string(name_list, ', ') from last_run
-- WHERE ARRAY[parent_id] = id_list
-- ORDER BY id_list;

WITH RECURSIVE last_run(parent_id, id_list, name_list) AS (
  SELECT parent_id, ARRAY[id] AS id_list, ARRAY[name] AS name_list FROM category 
  WHERE parent_id IS NULL
  UNION ALL
  SELECT category.parent_id, array_cat(ARRAY[category.id], last_run.id_list), array_cat(ARRAY[category.name], last_run.name_list) FROM last_run
  JOIN category ON last_run.id_list[1] = category.parent_id 
  ) 
SELECT id_list, array_to_string(name_list, ', ') AS name_list FROM last_run, category
WHERE last_run.name_list[1] = category.name
ORDER BY id_list;

ROLLBACK;

-- NOTES AND TRIAL/ERROR (should have had a timer going...⏰):

-- can't use array_prepend or array_append because they take a single value and 'push' or 'shift' onto the array - - must use array_cat to combine multiple array vals

-- must select all three values in the second SELECT so they are cast to arrays and concatenated

-- ERROR:  aggregate functions are not allowed in a recursive query's recursive term

-- The UNION operator combines result sets of two or more SELECT statements into a single result set. 

--  parent_id | id_list | name_list 
-- -----------+---------+-----------
--            |       1 | animal
--            |       2 | mineral
--            |       3 | vegetable
--          1 |       4 | dog
--          1 |       5 | cat
--          4 |       6 | doberman
--          4 |       7 | dachshund
--          3 |       8 | carrot
--          3 |       9 | lettuce
--         11 |      10 | paradox1
--         10 |      11 | paradox2

--  id_list |       name_list        
-- ---------+------------------------
--  {1}     | animal
--  {2}     | mineral
--  {3}     | vegetable
--  {4,1}   | dog, animal
--  {5,1}   | cat, animal
--  {6,4,1} | doberman, dog, animal
--  {7,4,1} | dachshund, dog, animal
--  {8,3}   | carrot, vegetable
--  {9,3}   | lettuce, vegetable
--  {10,11} | paradox1, paradox2
--  {11,10} | paradox2, paradox1
