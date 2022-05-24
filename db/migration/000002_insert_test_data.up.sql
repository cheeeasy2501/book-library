-- INSERT PUBLISH HOUSES
INSERT INTO house_publishes (id, name, created_at, updated_at)
VALUES (1, 'Test Publish House 1', now(), now());
INSERT INTO house_publishes (id, name, created_at, updated_at)
VALUES (2, 'Test Publish House 2', now(), now());
INSERT INTO house_publishes (id, name, created_at, updated_at)
VALUES (3, 'Test Publish House 3', now(), now());
INSERT INTO house_publishes (id, name, created_at, updated_at)
VALUES (4, 'Test Publish House 4', now(), now());
-- INSERT BOOKS
INSERT INTO books (id, publishhouse_id, title, description, link, in_stock, created_at, updated_at)
VALUES (1, 1, 'Test Book Title 1',
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam vulputate tellus metus, vel lobortis sapien egestas nec. Quisque non vestibulum mi. Proin malesuada lectus orci, quis tempor orci congue in. Cras sem elit, pretium vel posuere eu, luctus nec mi. Donec vulputate dolor in imperdiet volutpat. Ut sit amet ipsum ante. Vestibulum sapien nunc, sollicitudin at dictum eget, scelerisque nec lorem. Nulla bibendum rhoncus elit vel cursus. Sed feugiat iaculis nunc, quis efficitur odio vulputate et. Maecenas ut gravida massa. Duis arcu justo, venenatis eu laoreet et, placerat at est. Quisque varius ipsum dolor, aliquam fermentum orci ultrices condimentum. Duis quis dolor magna. Nunc aliquet purus sit amet ligula elementum vulputate. In placerat arcu placerat eros feugiat, sed lobortis justo sodales.',
        'http://test-1.com', 1, now(), now());

INSERT INTO books (id, publishhouse_id, title, description, link, in_stock, created_at, updated_at)
VALUES (2, 2, 'Test Book Title 2',
        'Duis lobortis molestie sagittis. Donec viverra enim ac posuere lobortis. Curabitur varius nec mauris sit amet facilisis. Etiam tincidunt egestas purus, sed fermentum nisl suscipit non. Sed dictum volutpat congue. Phasellus venenatis orci sem, ac laoreet mi lacinia sed. Proin vehicula diam diam, vel pretium ipsum lobortis ut. Donec condimentum fringilla risus, non rhoncus sapien ultricies pellentesque. Phasellus fringilla diam et consectetur venenatis. Mauris iaculis enim ut cursus consequat. Integer vitae magna accumsan, pretium odio sed, vulputate velit. Morbi mauris tortor, rutrum eu fermentum sit amet, hendrerit ut mi. Mauris venenatis, metus et mollis tincidunt, felis lectus lobortis augue, non mollis est nisi nec nisl. Proin quam lorem, finibus id mauris vel, eleifend ullamcorper arcu. Phasellus urna mauris, congue non felis non, porta sagittis mi. Quisque odio mi, pharetra sit amet lectus ut, accumsan finibus orci.',
        'http://test-2.com', 2, now(), now());

INSERT INTO books (id, publishhouse_id, title, description, link, in_stock, created_at, updated_at)
VALUES (3, 3, 'Test Book Title 3',
        'Donec auctor est a condimentum ultricies. Nunc vitae orci egestas, aliquet tellus vel, efficitur purus. Cras vel ex id lectus egestas finibus sed dictum mauris. Nam euismod elementum arcu sed malesuada. Nam cursus fermentum nunc, eget lobortis lorem dapibus ut. Curabitur quis eleifend erat. Vestibulum blandit cursus leo, eget volutpat tortor blandit vitae. Curabitur euismod iaculis tincidunt. Ut et aliquet nunc. Cras sit amet arcu ex. Phasellus tellus eros, eleifend nec ipsum eu, hendrerit scelerisque metus. Quisque egestas felis vitae orci sodales ultricies. Etiam molestie iaculis orci suscipit rhoncus. Nulla imperdiet dui non lacus commodo luctus. Nulla vehicula quis urna nec ultricies.',
        'http://test-3.com', 2, now(), now());

INSERT INTO books (id, publishhouse_id, title, description, link, in_stock, created_at, updated_at)
VALUES (4, 4, 'Test Book Title 4',
        'Nunc facilisis sapien et massa laoreet, nec tincidunt urna tristique. Aliquam vel tellus non arcu mattis vulputate. Cras ultrices molestie feugiat. Aenean iaculis pellentesque quam, non feugiat nulla rhoncus non. Sed fermentum vestibulum tempor. Nam nec lorem non ipsum tristique eleifend. Vivamus nec nisi rutrum, tempor lorem nec, sollicitudin tortor. In congue eleifend libero vitae dictum.',
        'http://test-4.com', 3, now(), now());

INSERT INTO books (id, publishhouse_id, title, description, link, in_stock, created_at, updated_at)
VALUES (5, 1, 'Test Book Title 5',
        'Quisque vel est malesuada, dapibus purus quis, fermentum eros. Proin quis mauris ligula. Vivamus nec maximus quam. Donec varius erat a lectus maximus tempus. Ut tempus dapibus dolor in gravida. Sed urna turpis, feugiat eget egestas rutrum, faucibus eget ante. Cras porta ipsum a nibh maximus commodo. Vivamus accumsan ornare dolor ut condimentum. Nulla congue egestas sapien, vitae ultricies tellus pharetra laoreet. Morbi id velit blandit, pulvinar ex sed, accumsan turpis. Etiam euismod auctor porta. Integer rutrum arcu ac erat consequat, at viverra massa interdum. Ut finibus mauris lacus, a facilisis nisi tempor sit amet. Vivamus sed turpis id arcu tempus pharetra. Phasellus eleifend nec turpis vel suscipit.',
        'http://test-5.com', 1, now(), now());

INSERT INTO books (id, publishhouse_id, title, description, link, in_stock, created_at, updated_at)
VALUES (6, 1, 'Test Book Title 6',
        'Etiam nec interdum enim, et tincidunt urna. Aliquam vel leo vitae mauris mattis mollis. Pellentesque eget facilisis massa. Aliquam posuere aliquet hendrerit. Mauris pulvinar nibh elit, scelerisque consectetur tortor facilisis ut. Praesent vehicula ut risus eget facilisis. Pellentesque ac massa in dolor ullamcorper eleifend id in ligula. Suspendisse vitae finibus ante. Vestibulum ut tellus tincidunt, venenatis est sit amet, aliquet ipsum. Suspendisse facilisis mi odio, vulputate vehicula enim luctus vel. Etiam sit amet lacus facilisis, luctus neque non, vehicula libero.',
        'http://test-6.com', 4, now(), now());

INSERT INTO books (id, publishhouse_id, title, description, link, in_stock, created_at, updated_at)
VALUES (7, 2, 'Test Book Title 7',
        'Maecenas convallis eget lacus ac sodales. Pellentesque eget sodales ligula. Donec sollicitudin luctus dolor mattis volutpat. Nullam semper elementum finibus. Vestibulum varius vitae mi in volutpat. Interdum et malesuada fames ac ante ipsum primis in faucibus. Integer nec ex ut turpis accumsan faucibus sit amet sit amet arcu. Sed scelerisque in sapien eget vehicula.',
        'http://test-7.com', 10, now(), now());

--INSERT AUTHORS
INSERT INTO author (id, firstname, lastname, created_at, updated_at)
VALUES (1, 'Inessa', 'Rhouben', now(), now());
INSERT INTO author (id, firstname, lastname, created_at, updated_at)
VALUES (2, 'Gisilbert', 'Brigid', now(), now());
INSERT INTO author (id, firstname, lastname, created_at, updated_at)
VALUES (3, 'Soroush', 'Wulfgifu', now(), now());
INSERT INTO author (id, firstname, lastname, created_at, updated_at)
VALUES (4, 'Anastasija', 'Fiachna', now(), now());
INSERT INTO author (id, firstname, lastname, created_at, updated_at)
VALUES (5, 'Gilbert', 'Martin', now(), now());
INSERT INTO author (id, firstname, lastname, created_at, updated_at)
VALUES (6, 'Viltautė', 'Lady', now(), now());
INSERT INTO author (id, firstname, lastname, created_at, updated_at)
VALUES (7, 'Connla', 'Jyoti', now(), now());
INSERT INTO author (id, firstname, lastname, created_at, updated_at)
VALUES (8, 'Hervey', 'Angel', now(), now());
INSERT INTO author (id, firstname, lastname, created_at, updated_at)
VALUES (9, 'Kenina', 'Ace', now(), now());
INSERT INTO author (id, firstname, lastname, created_at, updated_at)
VALUES (10, 'Jehan', 'Jannine', now(), now());


-- INSERT AUTHOR_BOOKS RELATIONS
INSERT INTO author_books (id, author_id, book_id)
VALUES (1, 1, 1);
INSERT INTO author_books (id, author_id, book_id)
VALUES (2, 1, 3);
INSERT INTO author_books (id, author_id, book_id)
VALUES (3, 1, 6);
INSERT INTO author_books (id, author_id, book_id)
VALUES (4, 2, 4);
INSERT INTO author_books (id, author_id, book_id)
VALUES (5, 2, 2);
INSERT INTO author_books (id, author_id, book_id)
VALUES (6, 2, 7);
INSERT INTO author_books (id, author_id, book_id)
VALUES (7, 3, 3);
INSERT INTO author_books (id, author_id, book_id)
VALUES (8, 3, 5);
INSERT INTO author_books (id, author_id, book_id)
VALUES (9, 4, 5);
INSERT INTO author_books (id, author_id, book_id)
VALUES (10, 4, 1);
INSERT INTO author_books (id, author_id, book_id)
VALUES (11, 4, 7);
INSERT INTO author_books (id, author_id, book_id)
VALUES (12, 5, 6);
INSERT INTO author_books (id, author_id, book_id)
VALUES (13, 6, 7);
INSERT INTO author_books (id, author_id, book_id)
VALUES (14, 6, 1);
INSERT INTO author_books (id, author_id, book_id)
VALUES (15, 7, 2);
INSERT INTO author_books (id, author_id, book_id)
VALUES (16, 8, 2);
INSERT INTO author_books (id, author_id, book_id)
VALUES (17, 9, 1);
INSERT INTO author_books (id, author_id, book_id)
VALUES (18, 10, 5);
INSERT INTO author_books (id, author_id, book_id)
VALUES (19, 10, 1);
INSERT INTO author_books (id, author_id, book_id)
VALUES (20, 10, 4);



