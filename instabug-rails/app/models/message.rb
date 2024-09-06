class Message < ApplicationRecord
  belongs_to :chat

  # include Elasticsearch::Model
  # include Elasticsearch::Model::Callbacks

  validates :body, presence: true
  validates :number, presence: true, uniqueness: { scope: :chat_id }

  before_validation :generate_number, on: :create

  # # Customize the JSON serialization for Elasticsearch
  # def as_indexed_json(options = {})
  #   as_json(only: [ :body, :number, :chat_id ])
  # end

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
