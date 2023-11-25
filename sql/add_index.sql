create index ng_words_user_id_livestream_id_created_at_index2
    on ng_words (user_id asc, livestream_id asc, created_at desc);

create index livestream_tags_livestream_id_index
        on livestream_tags (livestream_id);


create index reservation_slots_start_at_end_at_index
    on reservation_slots (start_at, end_at);

create index reservation_slots_start_at_index
    on reservation_slots (start_at);

create index reservation_slots_end_at_index
    on reservation_slots (end_at);

create index themes_user_id_index
    on themes (user_id);


create index livestream_tags_livestream_id_index2
    on livestream_tags (livestream_id);

create index reactions_created_at_livestream_id_index
    on reactions (created_at, livestream_id);


create index `id_user_id_idx_lvs`
    on livestreams (id, user_id);


create index livecomments_livestream_id_id_index
    on livecomments (livestream_id, id);


