



-- make sure you are connected to chat_app database
GRANT SELECT ON ALL TABLES IN SCHEMA public TO chat_app_ro;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO chat_app_ro;

GRANT USAGE ON ALL SEQUENCES IN SCHEMA public TO chat_app_ro;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT USAGE ON SEQUENCES TO chat_app_ro;
