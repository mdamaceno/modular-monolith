module Components
  module Order
    module Controller
      include Components::Order::UseCase
      include Helpers

      def self.extended(base)
        base.class_exec do
          post '/checkout' do
            Serializer.serialize(CreateOrder.new.call).to_json
          end
        end
      end
    end
  end
end
