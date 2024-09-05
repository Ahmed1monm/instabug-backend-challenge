class Message < ApplicationRecord
  belongs_to :chat

  validates :body, presence: true
  validates :number, presence: true, uniqueness: { scope: :chat_id }

  before_validation :generate_number, on: :create

  private

  def generate_number
    Message.transaction do
      max_number = self.chat.messages.lock(true).maximum(:number) || 0
      self.number = max_number + 1
    end
  rescue ActiveRecord::StatementInvalid => e
    retries ||= 0
    retries += 1
    retry if retries
  end

end
