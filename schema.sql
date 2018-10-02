CREATE TABLE mecha (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255),
    exts JSON
);

INSERT INTO mecha
(name, exts)
VALUES (
    'Aim for the top!',
    '{
        "mechs": ["Gunbuster", "Sizzler"],
        "director": "Hideaki Anno",
        "budget_woes": true,
        "space_monsters": 16000000000
    }'
);

INSERT INTO mecha
(name, exts)
VALUES (
    'SDF Macross',
    '{
        "mechs": ["namekyrie", "Defender"],
        "director": "Noboru Ishiguro",
        "budget_woes": true,
        "space_monsters": 0
    }'
);

INSERT INTO mecha
(name, exts)
VALUES (
    'Gundam Wing',
    '{
        "mechs": ["Deathscythe", "Sandrock"],
        "director": "Masashi Ikeda",
        "budget_woes": false,
        "space_monsters": 0
    }'
);

INSERT INTO mecha
(name, exts)
VALUES (
    'Neon Genesis Evangelion',
    '{
        "mechs": ["Unit-00", "Unit-01"],
        "director": "Hideaki Anno",
        "budget_woes": true,
        "space_monsters": 18
    }'
);
