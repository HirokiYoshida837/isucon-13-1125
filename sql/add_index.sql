-- create index ng_words_user_id_livestream_id_created_at_index2
--     on ng_words (user_id asc, livestream_id asc, created_at desc);


-- index 増やしたい場合はこんな感じで
set @x := (select count(*)
           from information_schema.statistics
           where table_name = 'ng_words'
             and index_name = 'ng_words_user_id_livestream_id_created_at_index2'
             and table_schema = database());
set @sql := if(@x > 0, 'select ''Index exists.''',
    'create index ng_words_user_id_livestream_id_created_at_index2
        on ng_words (user_id asc, livestream_id asc, created_at desc);');
PREPARE stmt FROM @sql;
EXECUTE stmt;



set @x := (select count(*)
           from information_schema.statistics
           where table_name = 'livestream_tags'
             and index_name = 'livestream_tags_livestream_id_index'
             and table_schema = database());
set @sql := if(@x > 0, 'select ''Index exists.''',
    'create index livestream_tags_livestream_id_index
        on livestream_tags (livestream_id);;');
PREPARE stmt FROM @sql;
EXECUTE stmt;
