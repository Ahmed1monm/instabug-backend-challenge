require "redis"

$redis = Redis.new(driver: :ruby, url: "redis://127.0.0.1:6379/1")
