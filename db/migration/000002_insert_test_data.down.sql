--DELETE AUTHOR_BOOKS RELATIONS
DELETE
FROM author_books
WHERE id <= 20;
--DELETE AUTHORS
DELETE
FROM author
WHERE id <= 10;
--DELETE PUBLISH HOUSES
DELETE
FROM house_publishes
WHERE id <= 4;
--DELETE BOOKS
DELETE
FROM books
WHERE id <= 7;



