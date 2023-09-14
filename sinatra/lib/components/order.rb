module Components
  module Order
    module Controller
      include Components::Order::UseCase
      include Helpers

      def self.extended(base)
        base.class_exec do
          post '' do
            Serializer.serialize(CreateOrder.new.call)
          end
        end
      end
    end
  end
end
