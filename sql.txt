PRAGMA foreign_keys = ON;

CREATE TABLE Material_type (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    material_type TEXT UNIQUE NOT NULL,
    loss_percentage REAL NOT NULL CHECK(loss_percentage >= 0 AND loss_percentage <= 100)
);

INSERT INTO Material_type (material_type, loss_percentage)
VALUES 
    ('Пластичные материалы', 0.12),
    ('Добавка', 0.20),
    ('Электролит', 0.15),
    ('Глазурь', 0.30),
    ('Пигмент', 0.25);

CREATE TABLE Materials (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    material_name TEXT UNIQUE NOT NULL,
    material_type_id INTEGER NOT NULL,
    unit_price REAL NOT NULL CHECK(unit_price > 0),
    stock_quantity REAL NOT NULL CHECK(stock_quantity >= 0),
    min_quantity REAL NOT NULL CHECK(min_quantity >= 0),
    package_quantity REAL NOT NULL CHECK(package_quantity > 0),
    unit_of_measure TEXT NOT NULL,
    FOREIGN KEY (material_type_id) REFERENCES Material_type(id)
        ON UPDATE CASCADE ON DELETE RESTRICT
);

INSERT INTO Materials (
    material_name, 
    material_type_id,
    unit_price,
    stock_quantity,
    min_quantity,
    package_quantity,
    unit_of_measure
) VALUES
('Глина', 1, 15.29, 1570.00, 5500.00, 30, 'кг'),
('Каолин', 1, 18.20, 1030.00, 3500.00, 25, 'кг'),
('Гидрослюда', 1, 17.20, 2147.00, 3500.00, 25, 'кг'),
('Монтмориллонит', 1, 17.67, 3000.00, 3000.00, 30, 'кг'),
('Перлит', 2, 13.99, 150.00, 1000.00, 50, 'л'),
('Стекло', 2, 2.40, 3000.00, 1500.00, 500, 'кг'),
('Дегидратированная глина', 2, 21.95, 3000.00, 2500.00, 20, 'кг'),
('Шамот', 2, 27.50, 2300.00, 1960.00, 20, 'кг'),
('Кварцевый песок', 2, 4.29, 3000.00, 1600.00, 50, 'кг'),
('Жильный кварц', 2, 18.60, 2556.00, 1600.00, 25, 'кг'),
('Техническая сода', 3, 54.55, 1200.00, 1500.00, 25, 'кг'),
('Жидкое стекло', 3, 76.59, 500.00, 1500.00, 15, 'кг'),
('Углещелочной реагент', 3, 3.45, 450.00, 1100.00, 25, 'кг'),
('Пирофосфат натрия', 3, 700.99, 356.00, 1200.00, 25, 'кг'),
('Кварц', 4, 375.96, 1500.00, 2500.00, 10, 'кг'),
('Полевой шпат', 4, 15.99, 750.00, 1500.00, 100, 'кг'),
('Барий углекислый', 4, 303.64, 340.00, 1500.00, 25, 'кг'),
('Бура техническая', 4, 125.99, 165.00, 1300.00, 25, 'кг'),
('Краска-раствор', 5, 200.90, 1496.00, 2500.00, 5, 'л'),
('Порошок цветной', 5, 84.39, 511.00, 1750.00, 25, 'кг');

CREATE TABLE Suppliers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    supplier_name TEXT UNIQUE NOT NULL,
    supplier_type TEXT NOT NULL,
    inn TEXT UNIQUE NOT NULL CHECK(length(inn) BETWEEN 10 AND 12),
    rating REAL,
    start_date TEXT NOT NULL CHECK(date(start_date) IS NOT NULL)
);

INSERT INTO Suppliers (
    supplier_name,
    supplier_type,
    inn,
    rating,
    start_date
) VALUES
    ('БрянскСтройресурс', 'ЗАО', '9432455179', 8, '2015-12-20'),
    ('Стройкомплект', 'ЗАО', '7803888520', 7, '2017-09-13'),
    ('Железногорская руда', 'ООО', '8430391035', 7, '2016-12-23'),
    ('Белая гора', 'ООО', '4318170454', 8, '2019-05-27'),
    ('Тульский обрабатывающий завод', 'ООО', '7687851800', 7, '2015-06-16'),
    ('ГорТехРазработка', 'ПАО', '6119144874', 9, '2021-12-27'),
    ('Сапфир', 'ОАО', '1122170258', 3, '2022-04-10'),
    ('ХимБытСервис', 'ПАО', '8355114917', 5, '2022-03-13'),
    ('ВоронежРудоКомбинат', 'ОАО', '3532367439', 8, '2023-11-11'),
    ('Смоленский добывающий комбинат', 'ОАО', '2362431140', 3, '2018-11-23'),
    ('МосКарьер', 'ПАО', '4159215346', 2, '2012-07-07'),
    ('КурскРесурс', 'ЗАО', '9032455179', 4, '2021-07-23'),
    ('Нижегородская разработка', 'ОАО', '3776671267', 9, '2016-05-23'),
    ('Речная долина', 'ОАО', '7447864518', 8, '2015-06-25'),
    ('Карелия добыча', 'ПАО', '9037040523', 6, '2017-03-09'),
    ('Московский ХимЗавод', 'ПАО', '6221520857', 4, '2015-05-07'),
    ('Горная компания', 'ЗАО', '2262431140', 3, '2020-12-22'),
    ('Минерал Ресурс', 'ООО', '4155215346', 7, '2015-05-22'),
    ('Арсенал', 'ЗАО', '3961234561', 5, '2010-11-25'),
    ('КамчаткаСтройМинералы', 'ЗАО', '9600275878', 7, '2016-12-20');

CREATE TABLE Material_suppliers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    material_id INTEGER NOT NULL,
    supplier_id INTEGER NOT NULL,
    UNIQUE (material_id, supplier_id),
    FOREIGN KEY (material_id) REFERENCES Materials(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (supplier_id) REFERENCES Suppliers(id)
        ON UPDATE CASCADE ON DELETE CASCADE
);

INSERT INTO Material_suppliers (material_id, supplier_id)
SELECT m.id AS material_id, s.id AS supplier_id
FROM 
    (SELECT 'Краска-раствор' AS material_name, 'Арсенал' AS supplier_name UNION ALL
     SELECT 'Каолин', 'Железногорская руда' UNION ALL
     SELECT 'Каолин', 'ВоронежРудоКомбинат' UNION ALL
     SELECT 'Стекло', 'Арсенал' UNION ALL
     SELECT 'Кварцевый песок', 'БрянскСтройресурс' UNION ALL
     SELECT 'Перлит', 'ГорТехРазработка' UNION ALL
     SELECT 'Глина', 'Белая гора' UNION ALL
     SELECT 'Кварцевый песок', 'КамчаткаСтройМинералы' UNION ALL
     SELECT 'Дегидратированная глина', 'ВоронежРудоКомбинат' UNION ALL
     SELECT 'Полевой шпат', 'Белая гора' UNION ALL
     SELECT 'Глина', 'БрянскСтройресурс' UNION ALL
     SELECT 'Порошок цветной', 'Арсенал' UNION ALL
     SELECT 'Жильный кварц', 'Горная компания' UNION ALL
     SELECT 'Полевой шпат', 'БрянскСтройресурс' UNION ALL
     SELECT 'Гидрослюда', 'ВоронежРудоКомбинат' UNION ALL
     SELECT 'Стекло', 'КамчаткаСтройМинералы' UNION ALL
     SELECT 'Полевой шпат', 'КурскРесурс' UNION ALL
     SELECT 'Монтмориллонит', 'ВоронежРудоКомбинат' UNION ALL
     SELECT 'Кварц', 'МосКарьер' UNION ALL
     SELECT 'Барий углекислый', 'Горная компания' UNION ALL
     SELECT 'Жидкое стекло', 'КурскРесурс' UNION ALL
     SELECT 'Шамот', 'Горная компания' UNION ALL
     SELECT 'Глина', 'ГорТехРазработка' UNION ALL
     SELECT 'Кварц', 'Карелия добыча' UNION ALL
     SELECT 'Гидрослюда', 'Железногорская руда' UNION ALL
     SELECT 'Перлит', 'ВоронежРудоКомбинат' UNION ALL
     SELECT 'Шамот', 'Арсенал' UNION ALL
     SELECT 'Барий углекислый', 'КамчаткаСтройМинералы' UNION ALL
     SELECT 'Бура техническая', 'КамчаткаСтройМинералы' UNION ALL
     SELECT 'Техническая сода', 'Минерал Ресурс' UNION ALL
     SELECT 'Пирофосфат натрия', 'КамчаткаСтройМинералы' UNION ALL
     SELECT 'Гидрослюда', 'Белая гора' UNION ALL
     SELECT 'Жильный кварц', 'Карелия добыча' UNION ALL
     SELECT 'Перлит', 'Смоленский добывающий комбинат' UNION ALL
     SELECT 'Кварцевый песок', 'Карелия добыча' UNION ALL
     SELECT 'Монтмориллонит', 'Белая гора' UNION ALL
     SELECT 'Краска-раствор', 'КурскРесурс' UNION ALL
     SELECT 'Стекло', 'Сапфир' UNION ALL
     SELECT 'Порошок цветной', 'КурскРесурс' UNION ALL
     SELECT 'Каолин', 'БрянскСтройресурс' UNION ALL
     SELECT 'Жидкое стекло', 'Минерал Ресурс' UNION ALL
     SELECT 'Бура техническая', 'Сапфир' UNION ALL
     SELECT 'Дегидратированная глина', 'МосКарьер' UNION ALL
     SELECT 'Бура техническая', 'Нижегородская разработка' UNION ALL
     SELECT 'Монтмориллонит', 'Железногорская руда' UNION ALL
     SELECT 'Жидкое стекло', 'Московский ХимЗавод' UNION ALL
     SELECT 'Жидкое стекло', 'Сапфир' UNION ALL
     SELECT 'Порошок цветной', 'Московский ХимЗавод' UNION ALL
     SELECT 'Порошок цветной', 'ХимБытСервис' UNION ALL
     SELECT 'Углещелочной реагент', 'Московский ХимЗавод' UNION ALL
     SELECT 'Кварц', 'Речная долина' UNION ALL
     SELECT 'Жильный кварц', 'Нижегородская разработка' UNION ALL
     SELECT 'Краска-раствор', 'Московский ХимЗавод' UNION ALL
     SELECT 'Кварц', 'Нижегородская разработка' UNION ALL
     SELECT 'Дегидратированная глина', 'Стройкомплект' UNION ALL
     SELECT 'Кварцевый песок', 'Речная долина' UNION ALL
     SELECT 'Барий углекислый', 'Сапфир' UNION ALL
     SELECT 'Шамот', 'Стройкомплект' UNION ALL
     SELECT 'Дегидратированная глина', 'Сапфир' UNION ALL
     SELECT 'Гидрослюда', 'Тульский обрабатывающий завод' UNION ALL
     SELECT 'Пирофосфат натрия', 'Сапфир' UNION ALL
     SELECT 'Перлит', 'Тульский обрабатывающий завод' UNION ALL
     SELECT 'Техническая сода', 'Сапфир' UNION ALL
     SELECT 'Глина', 'Смоленский добывающий комбинат' UNION ALL
     SELECT 'Техническая сода', 'Московский ХимЗавод' UNION ALL
     SELECT 'Монтмориллонит', 'Смоленский добывающий комбинат' UNION ALL
     SELECT 'Углещелочной реагент', 'КурскРесурс' UNION ALL
     SELECT 'Техническая сода', 'ХимБытСервис' UNION ALL
     SELECT 'Бура техническая', 'Стройкомплект' UNION ALL
     SELECT 'Пирофосфат натрия', 'Тульский обрабатывающий завод' UNION ALL
     SELECT 'Жильный кварц', 'Смоленский добывающий комбинат' UNION ALL
     SELECT 'Стекло', 'Стройкомплект' UNION ALL
     SELECT 'Углещелочной реагент', 'ХимБытСервис' UNION ALL
     SELECT 'Барий углекислый', 'Тульский обрабатывающий завод' UNION ALL
     SELECT 'Пирофосфат натрия', 'ХимБытСервис' UNION ALL
     SELECT 'Каолин', 'Тульский обрабатывающий завод' UNION ALL
     SELECT 'Шамот', 'МосКарьер' UNION ALL
     SELECT 'Углещелочной реагент', 'Тульский обрабатывающий завод' UNION ALL
     SELECT 'Краска-раствор', 'ХимБытСервис' UNION ALL
     SELECT 'Полевой шпат', 'Смоленский добывающий комбинат'
    ) AS pairs
JOIN Materials m ON m.material_name = pairs.material_name
JOIN Suppliers s ON s.supplier_name = pairs.supplier_name;

CREATE TABLE Product_type (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    product_type TEXT UNIQUE NOT NULL,
    type_coefficient REAL NOT NULL CHECK(type_coefficient > 0)
);

INSERT INTO Product_type (id, product_type, type_coefficient) VALUES
(1, 'Тип продукции 1', 1.2),
(2, 'Тип продукции 2', 8.59),
(3, 'Тип продукции 3', 3.45),
(4, 'Тип продукции 4', 5.6);