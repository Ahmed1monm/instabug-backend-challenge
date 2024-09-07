class RenameCounterCacheColumns < ActiveRecord::Migration[7.2]
  def change
    rename_column :applications, :chats_number, :chats_count
    rename_column :chats, :messages_number, :messages_count
  end
end
