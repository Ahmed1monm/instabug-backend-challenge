class AddTokenToApplications < ActiveRecord::Migration[7.2]
  def change
    add_index :applications, :token
  end
end
