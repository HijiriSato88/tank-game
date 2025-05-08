-- イベント作成コマンド

INSERT INTO events (name, slug, start_at, end_at)
VALUES (
  '1年間のテストイベント',
  'test_event_yearly',
  NOW(),
  DATE_ADD(NOW(), INTERVAL 1 YEAR)
);
