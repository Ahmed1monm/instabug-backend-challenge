class AddChatsNumberToApplication < ActiveRecord::Migration[7.2]
  def change
    add_column :applications, :chats_number, :integer
    add_column :chats, :messages_number, :integer
  end
end
