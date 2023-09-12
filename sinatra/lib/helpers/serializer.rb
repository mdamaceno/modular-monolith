module Helpers
  module Serializer
    def self.serialize(hash)
      { data: hash }.to_json
    end
  end
end
