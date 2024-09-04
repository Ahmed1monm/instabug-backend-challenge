class CreateMessage < ActiveRecord::Migration[7.2]
  def change
    create_table :messages do |t|
      t.string :body
      t.integer :number

      t.timestamps
    end
  end
end
