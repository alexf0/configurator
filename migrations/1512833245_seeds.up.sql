INSERT INTO configs (type, data, params) VALUES
('Develop.mr_robot','Database.processing','{"host": "localhost","port": "5432","database": "devdb","user": "mr_robot","password": "secret","schema": "public"}'),
('Test.vpn','Rabbit.log','{"host": "10.0.5.42","port": "5671","virtualhost": "/","user": "guest","password": "guest"}'),
('Develop.mr_robot', 'Redis.processing', '{"host":"localhost","port":"6379"}'),
('Production.user_prod', 'Database.db1', '{"host":"localhost","port":"5432","user":"postgres","password":"password","dbname":"db1","sslmode":"disable"}'),
('Develop.app_1','SSO.provider','{"url":"https://sso.example.ru/blitz/saml/profile/Metadata/SAML","app_id":"app_1"}')
ON CONFLICT (type, data) DO NOTHING;