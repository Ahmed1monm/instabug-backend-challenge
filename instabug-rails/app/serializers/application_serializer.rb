class ApplicationSerializer < ActiveModel::Serializer
  attributes :name, :token, :created_at, :updated_at
end
