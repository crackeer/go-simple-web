CREATE TABLE content(
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   title TEXT NOT NULL,
   content TEXT NOT NULL,
   category TEXT NOT NULL,
   content_type TEXT NOT NULL,
   username TEXT NOT NULL,
   create_at INTEGER DEFAULT '0',
   modify_at INTEGER DEFAULT '0'
);

CREATE TABLE share(
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   `table` TEXT NOT NULL,
   data_id TEXT NOT NULL,
   code TEXT NOT NULL,
   create_at INTEGER DEFAULT '0'
);