module Components
  module Marketing
    module Controller
      include Components::Marketing::UseCase
      include Helpers

      def self.extended(base)
        base.class_exec do
          get '/announces' do
            Serializer.serialize(AnnounceSomething.new.call)
          end

          post '/announces' do
            Serializer.serialize(AnnounceSomething.new.call)
          end
        end
      end
    end
  end
end
