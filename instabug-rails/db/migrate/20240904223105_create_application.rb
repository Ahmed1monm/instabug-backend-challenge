class CreateApplication < ActiveRecord::Migration[7.2]
  def change
    create_table :applications do |t|
      t.string :name
      t.string :token

      t.timestamps
    end
    create_table :messages do |t|
      t.string :body
      t.integer :number

      t.timestamps
    end
    create_table :chats do |t|
      t.string :name
      t.integer :number

      t.timestamps
    end
    add_index :applications, :token
    add_reference :messages, :chat, null: false, foreign_key: true
    add_reference :chats, :application, null: false, foreign_key: true
  end
end
