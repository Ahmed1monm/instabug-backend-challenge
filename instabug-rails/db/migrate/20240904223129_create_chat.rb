class CreateChat < ActiveRecord::Migration[7.2]
  def change
    create_table :chats do |t|
      t.string :name
      t.integer :number

      t.timestamps
    end
  end
end
