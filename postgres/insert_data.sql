INSERT INTO
    app_data.restaurants (name, menu)
VALUES
    (
        'Japanese Sushi',
        '[{"id": 1, "price": 15.99, "name": "Salmon sashimi"}, {"id": 2, "price": 12.50, "name": "California roll"}]'::jsonb
    ),
    (
        'Indian Curry House',
        '[{"id": 1, "price": 10.99, "name": "Chicken tikka masala"}, {"id": 2, "price": 9.50, "name": "Vegetable korma"}]'::jsonb
    ),
    (
        'Mexican Fiesta',
        '[{"id": 1, "price": 11.00, "name": "Tacos al pastor"}, {"id": 2, "price": 12.75, "name": "Chili con carne"}]'::jsonb
    );
    

INSERT INTO
    app_data.riders (name)
VALUES
    ('James'),
    ('Ben'),
    ('Timmy'),
    ('Anna');