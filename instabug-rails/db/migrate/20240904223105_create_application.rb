class CreateApplication < ActiveRecord::Migration[7.2]
  def change
    create_table :applications do |t|
      t.string :name
      t.string :token

      t.timestamps
    end
  end
end
