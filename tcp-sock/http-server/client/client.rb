require 'socket'

n=10
waitSec=3600

n.times do
  socket = Socket.new(:INET, :STREAM)
  fork do
    remote_addr = Socket.pack_sockaddr_in(8080, '192.168.10.3')
    socket.connect(remote_addr)
    sleep(waitSec)
  end
end
