class Message < ApplicationRecord
  include Elasticsearch::Model
  include Elasticsearch::Model::Callbacks

  belongs_to :chat

  validates :body, presence: true
  validates :number, presence: true, uniqueness: { scope: :chat_id }

  before_validation :generate_number, on: :create

  def self.search_in_chat(chat_id, query)
    search({
      query: {
        bool: {
          must: [
            { match: { body: query } },
            { term: { chat_id: chat_id } }
          ]
        }
      }
    })
  end

  settings index: { number_of_shards: 1 } do
    mappings dynamic: "false" do
      indexes :body, type: "text", analyzer: "english"
      indexes :chat_id, type: "keyword"
    end
  end

  def as_indexed_json(options = {})
    as_json(only: [:body, :chat_id, :number])
  end

  private

  def generate_number
    Message.transaction do
      max_number = self.chat.messages.lock(true).maximum(:number) || 0
      self.number = max_number + 1
    end
  rescue ActiveRecord::StatementInvalid => e
    retries ||= 0
    retries += 1
    retry if retries < 3
  end
end
