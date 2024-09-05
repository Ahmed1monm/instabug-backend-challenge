class Chat < ApplicationRecord
  belongs_to :application
  has_many :messages, dependent: :destroy

  before_validation :generate_number, on: :create
  validates :number, presence: true, uniqueness: { scope: :application_id }

  private

  def generate_number
    Chat.transaction do
      # Lock chats table to ensure the number is unique and sequential
      max_number = self.application.chats.lock(true).maximum(:number) || 0
      self.number = max_number + 1
    end
  rescue ActiveRecord::StatementInvalid => e
    # Handle potential deadlock or lock wait timeout
    retries ||= 0
    retries += 1
    retry if retries < 3
    raise e
  end
end
