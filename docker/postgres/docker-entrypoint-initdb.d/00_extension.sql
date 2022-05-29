--
-- DB に必要な関数や拡張モジュールを有効化
--

-- ULID の生成に必要な拡張機能を有効化する
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- ULID 生成用の関数を作成
DROP FUNCTION IF EXISTS gen_random_ulid;

CREATE FUNCTION gen_random_ulid() RETURNS UUID AS
$$
SELECT
  SUBSTRING(
      (
          SUBSTRING(
              DECODE(
                  TO_HEX(
                        int8 '4611686018427387904' +
                        ( -- b'0100000000000000000000000000000000000000000000000000000000000000'::bigint -- to_hexしたときに上位桁が0だと破棄される。必要な桁まで削除されてしまうのを防ぐ。
                            (
                                EXTRACT(EPOCH FROM NOW()) -- create epoch time
                                * 1000 -- ミリ秒化
                              )::BIGINT::BIT(64) -- マイクロ秒を切り捨て
                            & int8 '281474976710655'::BIT(64) -- 下位48bitのみを残すため。10889年まで枯渇しないので、今のところなくても良い
                          )::BIGINT
                    )
                , 'hex') -- byteaで扱う
            , 3) -- 上位16bitを削除
          || gen_random_bytes(10) -- ulidの下位80bitの生成
        )::TEXT
    , 3)::UUID; -- bytea -> textしたときの\xを削除しuuidにキャスト
$$ LANGUAGE SQL;
