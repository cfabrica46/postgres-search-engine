DROP TABLE IF EXISTS products;

CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(36) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    stock SERIAL NOT NULL,
    type VARCHAR(10) NOT NULL,
    brand VARCHAR(15) NOT NULL,
    price SERIAL NOT NULL,
    date DATE NOT NULL DEFAULT CURRENT_DATE,
    description VARCHAR(100) NOT NULL
);

INSERT INTO products(id,name,stock,type,brand,price,date,description)
    VALUES
        ('95fafff7-e70f-493c-9421-10b4eafd0b00',    'Apple Iphone 12 Mini',                         2,  'mobile',   'Apple',    960,    '2021-08-03',   'apple most optimal mobile'),
        ('f3d574e2-dc97-4d51-a064-e1f5d6841486',    'Apple iPhone 12',                              1,  'mobile',   'Apple',    1190,   '2021-08-04',   'apple most expensive mobile'),
        ('7edc909c-7e9c-473e-86ed-9c63dd743e2d',    'Apple Iphone 11',                              3,  'mobile',   'Apple',    840,    '2021-08-05',   'apple most cheapest mobile'),
        ('8e9f8668-3f80-4554-9620-c93a31b5210b',    'Galaxy A12',                                   2,  'mobile',   'Samsung',  165,    '2021-08-06',   'samsung most cheapest mobile'),
        ('453d21ce-03b0-4ea5-8668-8f5496d2f58f',    'Galaxy A21s',                                  2,  'mobile',   'Samsung',  195,    '2021-08-07',   'samsung most expensive mobile'),
        ('ef5ea118-669d-4355-9a6b-1187d2786099',    'Pavilion 13-bb0502la Intel Core i5-1135G7 ',   6,  'laptop',   'HP',       840,    '2021-08-08',   'hp most expensive laptop'),
        ('b9972450-0865-4761-9e43-8e0ca9a89b73',    'Jet Black 15T-DW300 15.6" Intel Core I5',      2,  'laptop',   'HP',       790,    '2021-08-09',   'hp most cheapest laptop'),
        ('4e802370-d70f-4018-bf56-ba17d8783cf0',    'IdeaPad Gaming i3',                            2,  'laptop',   'Lenovo',   1100,   '2021-08-10',   'lenovo most optimal laptop'),
        ('9bd83931-ba35-4b94-a4c5-5341eb78e25b',    'Chromebook 4',                                 4,  'laptop',   'Samsung',  255,    '2021-08-11',   'samsung most optimal laptop'),
        ('a46bbf2c-eae2-4bfb-8473-e6a4932a6cb5',    'Macbook Air 13',                               2,  'laptop',   'Apple',    1325,   '2021-08-12',   'apple most cheapest laptop'),
        ('0c1b7335-b255-4ea7-aa9b-f6a9af9066f8',    'Macbook Apple Pro 13',                         5,  'laptop',   'Apple',    1680,   '2021-08-13',   'apple most expensive laptop')
