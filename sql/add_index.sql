-- index 増やしたい場合はこんな感じで
create index ng_words_user_id_livestream_id_created_at_index2
    on ng_words (user_id asc, livestream_id asc, created_at desc);

create index livestream_tags_livestream_id_index
        on livestream_tags (livestream_id);



