class AddApplicationRefToChats < ActiveRecord::Migration[7.2]
  def change
    add_reference :chats, :application, null: false, foreign_key: true
  end
end
