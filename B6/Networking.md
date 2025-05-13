### Networking Basics
#### OSI model
- OSI stands for Open Systems Interconnection. It is a reference model that specifies standards for communications protocols and also the functionalities of each layer. 
##### Physical layer
- Purpose:
   - The Physical Layer is responsible for transmitting raw bits (0s and 1s) over a physical medium (like cables, fiber optics, or radio waves).
- Key Responsibilities:
   - Defining how bits are encoded into signals (electrical voltages, light pulses, or radio signals).
   - Specifying the physical characteristics: cables, connectors, signal levels, transmission rates (e.g., how many bits per second).
   - Handling synchronization of bits (making sure sender and receiver are synchronized on where a bit starts and ends).
   - Managing topology (how devices are physically connected) and transmission mode (simplex, half-duplex, full-duplex).

##### Data Link Layer (MAC Address & Switching)
- Purpose:
   - The Data-Link Layer is responsible for reliable transmission of data across a physical link. It takes the raw bits and organizes them into frames, detects and possibly corrects errors, and manages access to the medium.
- Key Responsibilities:
   - Framing: Grouping bits into structured units called frames (e.g., a packet with headers and data).
   - Addressing: Using MAC addresses to identify devices on the same network.
   - Error detection and correction: Using checksums like CRC (Cyclic Redundancy Check) to detect if frames were damaged.
   - Flow control: Preventing a fast sender from overwhelming a slow receiver.
   - Medium Access Control (MAC): Deciding who gets to use the physical medium when multiple devices want to send data (especially important in shared mediums like Wi-Fi).

##### Network Layer (IP Address & Routing)
- Purpose:
   - The Network Layer is responsible for moving packets between different networks. It decides the best path for the data to reach its destination (routing).
- Key Responsibilities:
   - Logical addressing (IP addresses: who is talking to whom across the world)
   - Routing (choosing the best path across the internet)
   - Packet forwarding (routers move packets based on destination IP)
   - Fragmentation (breaking up large packets if needed for smaller data links)
- Examples of Network Layer technologies:
   - IP (Internet Protocol, both IPv4 and IPv6)
   - Routers
   - ICMP (used for tools like ping)
<details>
<summary> Read more about: Network Layer</summary>

### Assigning Logical Address
- Logical addressing is the process of assigning unique IP addresses (IPv4 or IPv6) to devices within a network. Unlike physical addresses (MAC addresses), logical addresses can change based on network configurations. These addresses are hierarchical and help identify both the network and the device within that network. Logical addressing is important for:
  -  Enabling communication between devices on different networks.
  -  Facilitating routing by providing location-based information.
### Packetizing
- The process of encapsulating the data received from the upper layers of the network (also called payload) in a network layer packet at the source and decapsulating the payload from the network layer packet at the destination is known as packetizing. 
- The source host adds a header that contains the source and destination address and some other relevant information required by the network layer protocol to the payload received from the upper layer protocol and delivers the packet to the data link layer. 
- The destination host receives the network layer packet from its data link layer, decapsulates the packet, and delivers the payload to the corresponding upper layer protocol. The routers in the path are not allowed to change either the source or the destination address. The routers in the path are not allowed to decapsulate the packets they receive unless they need to be fragmented.  

### Host-to-Host Delivery
- The network layer ensures data is transferred from the source device (host) to the destination device (host) across one or multiple networks. This involves:
   - Determining the destination address.
   - Ensuring that data is transmitted without duplication or corruption.
   - Host-to-host delivery is a foundational aspect of communication in large-scale, interconnected systems like the internet.
### Forwarding
- Forwarding is the process of transferring packets between network devices such as routers, which are responsible for directing the packets toward their destination. When a router receives a packet from one of its attached networks, it needs to forward the packet to another attached network (unicast routing) or to some attached networks (in the case of multicast routing).The router uses:
   - Routing tables: These tables store information about possible paths to different networks.
   - Forwarding decisions: Based on the destination IP address in the packet header. Forwarding ensures that packets move closer to their destination efficiently.
### Fragmentation and Reassembly of Packets
- Some networks have a maximum transmission unit (MTU) that defines the largest packet size they can handle. If a packet exceeds the MTU, the network layer:
   - Fragments the packet into smaller pieces.
   - Adds headers to each fragment for identification and sequencing. At the destination, the fragments are reassembled into the original packet. This ensures compatibility with networks of varying capabilities without data loss.
   - Read more about Fragmentation at Network Layer.
### Logical Subnetting
- Logical subnetting involves dividing a large IP network into smaller, more manageable sub-networks (subnets). Subnetting helps:
   - Improve network performance by reducing congestion.
   - Enhance security by isolating parts of a network.
   - Simplify network management and troubleshooting. Subnetting uses subnet masks to define the range of IP addresses within each subnet, enabling efficient address allocation and routing.
### Network Address Translation (NAT)
- NAT allows multiple devices in a private network to share a single public IP address for internet access. This is achieved by:
   - Translating private IP addresses to a public IP address for outbound traffic.
   - Reversing the process for inbound traffic. Benefits of NAT include:
   - Conserving IPv4 addresses by reducing the need for unique public IPs for each device.
   - Enhancing security by masking internal IP addresses from external networks.
### Routing
- Routing is the process of moving data from one device to another device. These are two other services offered by the network layer. In a network, there are a number of routes available from the source to the destination. The network layer specifies some strategies which find out the best possible route. This process is referred to as routing. There are a number of routing protocols that are used in this process and they should be run to help the routers coordinate with each other and help in establishing communication throughout the network.
#### What is an IP Address?
- Imagine every device on the internet as a house. For you to send a letter to a friend living in one of these houses, you need their home address. In the digital world, this home address is what we call an IP (Internet Protocol) Address. It’s a unique string of numbers separated by periods (IPv4) or colons (IPv6) that identifies each device connected to the internet or a local network.
1. Based on Addressing Scheme (IPv4 vs. IPv6)
- IPv4 is the most common form of IP Address. It consists of four sets of numbers separated by dots. For example, 192.158.1.38. Each set of numbers can range from 0 to 255. This format can support over 4 billion unique addresses.
   - Example of IPv4 Address: 192.168.1.1
- IPv6 addresses were created to deal with the shortage of IPv4 addresses. They use 128 bits instead of 32, offering a vastly greater number of possible addresses. These addresses are expressed as eight groups of four hexadecimal digits, each group representing 16 bits. The groups are separated by colons.
   - Example of IPv6 Address: 2001:0db8:85a3:0000:0000:8a2e:0370:7334
   
2. Based on Usage (Public vs. Private)
- A Public IP address is assigned to every device that directly accesses the internet. This address is unique across the entire internet. Here are the key characteristics and uses of public IP addresses:
   - Uniqueness: Each public IP address is globally unique. No two devices on the internet can have the same public IP address at the same time.
   - Accessibility: Devices with a public IP address can be accessed directly from anywhere on the internet, assuming no firewall or security settings block the access.
   - Assigned by ISPs: Public IP addresses are assigned by Internet Service Providers (ISPs). When you connect to the internet through an ISP, your device or router receives a public IP address.
   - Types: Public IP addresses can be static (permanently assigned to a device) or dynamic (temporarily assigned and can change over time).
- Private IP addresses are used within private networks (such as home networks, office networks, etc.) and are not routable on the internet. This means that devices with private IP addresses cannot directly communicate with devices on the internet without a translating mechanism like a router performing Network Address Translation (NAT). Key features include:
   - Not globally unique: Private IP addresses are only required to be unique within their own network. Different private networks can use the same range of IP addresses without conflict.
   - Local communication: These addresses are used for communication between devices within the same network. They cannot be used to communicate directly with devices on the internet.
   - Defined ranges: The Internet Assigned Numbers Authority (IANA) has reserved specific IP address ranges for private use

3. Based on Assignment Method (Static vs. Dynamic)
- Static IP Addresses:
   - These are permanently assigned to a device, typically important for servers or devices that need a constant address.
   - Reliable for network services that require regular access such as websites, remote management.
- Dynamic IP Addresses:
   - Temporarily assigned from a pool of available addresses by the Dynamic Host Configuration Protocol (DHCP).
   - Cost-effective and efficient for providers, perfect for consumer devices that do not require permanent addresses.
#### Real World Scenario: Sending an Email from New York to Tokyo
Let’s explore how IP addresses work through a real-world example that involves sending an email from one person to another across the globe:
- Step 1: Assigning IP Addresses
   - Alice in New York wants to send an email to Bob in Tokyo.
   - Alice’s laptop has a private IP address (e.g., 192.168.1.5) assigned by her router at home.
Bob’s computer in Tokyo has a private IP address (e.g., 192.168.2.4) assigned by his router at his office.
- Step 2: Connection to the Internet
   - Both Alice and Bob’s routers have public IP addresses assigned by their Internet Service Providers (ISPs). These public IP addresses are what the devices use to send and receive data over the internet.
- Step 3: Sending the Email
   - Alice writes her email and hits send.
Her email service (e.g., Gmail) packages the message and its attachments into data packets. Each packet includes the source IP (Alice’s router’s public IP) and the destination IP (Bob’s email server’s public IP).
- Step 4: Routing the Packets
   - The data packets leave Alice’s laptop and travel to her home router. The router notes that the destination IP is outside the local network.
   - he router sends the packets to Alice’s ISP. The ISP uses routers that examine the destination IP address of the packets and determine the best route to send them toward their destination.
   - The packets may pass through several routers around the world — in data centers in countries like Canada, Germany, and finally Japan. Each router along the way reads the destination IP and forwards the packets accordingly.
- Step 5: Reaching Bob
   - The packets arrive at Bob’s email server’s ISP in Tokyo and are then forwarded to the server.
   - Bob’s email server reassembles the packets into the original email message.
- Step 6: Bob Accesses the Email
   - Bob’s computer requests the email from his server using his local network IP.
   - The server sends the email to Bob’s computer, allowing him to read the message Alice sent.
- Additional Details
   - NAT (Network Address Translation): Both Alice and Bob’s routers perform NAT, translating the private IP addresses to and from the public IP addresses when interfacing with the internet. This process is crucial for keeping the number of public IPs needed lower and adds a layer of security by masking internal network structures.
   
</details>

##### Transport Layer (TCP & UDP)
- Purpose:
   - The Transport Layer ensures that data is delivered reliably and in order between two devices. It also handles flow control and error recovery.
- Key Responsibilities:
   - Segmentation: Breaking large messages into smaller pieces (segments)
   - Reliable delivery: Using acknowledgments and retransmissions (TCP)
   - Multiplexing: Allowing multiple apps to use the network at once (via port numbers)
   - Flow control: Preventing a sender from overwhelming a receiver
- Examples of Transport Layer technologies:
   - TCP (Transmission Control Protocol — reliable)
   - UDP (User Datagram Protocol — faster but unreliable)
<details>
<summary> Read more about: TCP/UDP</summary>

### TCP
- Transmission Control Protocol (TCP) is a connection-oriented protocol for communications that helps in the exchange of messages between different devices over a network. It is one of the main protocols of the TCP/IP suite.
   - TCP establishes a reliable connection between sender and receiver using the three-way handshake (SYN, SYN-ACK, ACK) and it uses a four-step handshake (FIN, ACK, FIN, ACK) to close connections properly.
   - It ensures error-free, in-order delivery of data packets.
   - It uses acknowledgments (ACKs) to confirm receipt.
   - It prevents data overflow by adjusting the data transmission rate according to the receiver’s buffer size.
   - It prevents network congestion using algorithms like Slow Start, Congestion Avoidance, Fast Retransmit, and Fast Recovery.
   - TCP header uses checksum to detect corrupted data and requests retransmission if needed.
   - It is used in applications requiring reliable and ordered data transfer, such as web browsing, email, and remote login.
- Internet Protocol (IP) is a method that is useful for sending data from one device to another from all over the internet. It is a set of rules governing how data is sent and received over the internet. It is responsible for addressing and routing packets of data so they can travel from the sender to the correct destination across multiple networks. Every device contains a unique IP Address that helps it communicate and exchange data across other devices present on the internet.
#### Disadvantages of TCP
- TCP is made for Wide Area Networks, thus its size can become an issue for small networks with low resources. TCP runs several layers so it can slow down the speed of the network.
- It is not generic in nature. It cannot represent any protocol stack other than the TCP/IP suite. E.g., it cannot work with a Bluetooth connection.
- No modifications since their development around 30 years ago.

### UDP
- User Datagram Protocol (UDP) is a Transport Layer protocol. UDP is a part of the Internet Protocol suite, referred to as UDP/IP suite. Unlike TCP, it is an unreliable and connectionless protocol. So, there is no need to establish a connection before data transfer. The UDP helps to establish low-latency and loss-tolerating connections over the network. The UDP enables process-to-process communication.
#### Applications of UDP
- Used for simple request-response communication when the size of data is less and hence there is lesser concern about flow and error control.
- It is a suitable protocol for multicasting as UDP supports packet switching.
- UDP is used for some routing update protocols like RIP(Routing Information Protocol).
- Normally used for real-time applications which can not tolerate uneven delays between sections of a received message.
- VoIP (Voice over Internet Protocol) services, such as Skype and WhatsApp, use UDP for real-time voice communication. The delay in voice communication can be noticeable if packets are delayed due to congestion control, so UDP is used to ensure fast and efficient data transmission.
- DNS (Domain Name System) also uses UDP for its query/response messages. DNS queries are typically small and require a quick response time, making UDP a suitable protocol for this application.
- DHCP (Dynamic Host Configuration Protocol) uses UDP to dynamically assign IP addresses to devices on a network. DHCP messages are typically small, and the delay caused by packet loss or retransmission is generally not critical for this application.
#### How is UDP used in DDoS attacks?
A UDP flood attack is a type of Distributed Denial of Service (DDoS) attack where an attacker sends a large number of User Datagram Protocol (UDP) packets to a target port.
- UDP Protocol : Unlike TCP, UDP is connectionless and doesn’t require a handshake before data transfer. When a UDP packet arrives at a server, it checks the specified port for listening applications. If no app is found, the server sends an ICMP “destination unreachable” packet to the supposed sender (usually a random bystander due to spoofed IP addresses).
- Attack Process:
   - The attacker sends UDP packets with spoofed IP sender addresses to random ports on the target system.
   - The server checks each incoming packet’s port for a listening application (usually not found due to random port selection).
   - The server sends ICMP “destination unreachable” packets to the spoofed sender (random bystanders).
   - The attacker floods the victim with UDP data packets, overwhelming its resources.
- Mitigation : To protect against UDP flood attacks, monitoring network traffic for sudden spikes and implementing security measures are crucial. Organizations often use specialized tools and services to detect and mitigate such attacks effectively.

</details>

##### Session layer
- Purpose:
   - The Session Layer manages sessions: setting up, controlling, and ending communication sessions between two devices.
- Key Responsibilities:
   - Session establishment, maintenance, and termination
   - Synchronization: Managing checkpoints (useful for large data transfers or recoverable sessions)
   - Dialog control: Which side can send data and when (simplex/half-duplex/full-duplex)
- Examples of Session Layer technologies:
   - NetBIOS (Network Basic Input/Output System)
   - RPC (Remote Procedure Call)

<details>
<summary> Read more about: RPC</summary>

### What is Remote Procedure Call (RPC)?
- Remote Procedure Call (RPC) is a type of technology used in computing to enable a program to request a service from software located on another computer in a network without needing to understand the network’s details. RPC abstracts the complexities of the network by allowing the developer to think in terms of function calls rather than network details, facilitating the process of making a piece of software distributed across different systems.

- RPC works by allowing one program (a client) to directly call procedures (functions) on another machine (the server). The client makes a procedure call that appears to be local but is run on a remote machine. When an RPC is made, the calling arguments are packaged and transmitted across the network to the server. The server unpacks the arguments, performs the desired procedure, and sends the results back to the client.
### Working of a RPC
![plot](../img/operating-system-remote-call-procedure-working.png) 
1. A client invokes a client stub procedure, passing parameters in the usual way. The client stub resides within the client’s own address space. 

2. The client stub marshalls(pack) the parameters into a message. Marshalling includes converting the representation of the parameters into a standard format, and copying each parameter into the message. 

3. The client stub passes the message to the transport layer, which sends it to the remote server machine.  On the server, the transport layer passes the message to a server stub, which demarshalls(unpack) the parameters and calls the desired server routine using the regular procedure call mechanism. 

4. When the server procedure completes, it returns to the server stub (e.g., via a normal procedure call return), which marshalls the return values into a message.

5. The server stub then hands the message to the transport layer. The transport layer sends the result message back to the client transport layer, which hands the message back to the client stub. 

6. The client stub demarshalls the return parameters and execution returns to the caller.
### Types of RPC
- Callback RPC: Callback RPC allows processes to act as both clients and servers. It helps with remote processing of interactive applications. The server gets a handle to the client, and the client waits during the callback. This type of RPC manages callback deadlocks and enables peer-to-peer communication between processes.
- Broadcast RPC: In Broadcast RPC, a client’s request is sent to all servers on the network that can handle it. This type of RPC lets you specify that a client’s message should be broadcast. You can set up special broadcast ports. Broadcast RPC helps reduce the load on the network.
- Batch-mode RPC: Batch-mode RPC collects multiple RPC requests on the client side and sends them to the server in one batch. This reduces the overhead of sending many separate requests. Batch-mode RPC works best for applications that don’t need to make calls very often. It requires a reliable way to send data.

</details>

##### Presentation Layer
- Purpose:
   - The Presentation Layer deals with the syntax and semantics of the information being exchanged. It's where data is translated between formats, encrypted, or compressed.
- Key Responsibilities:
   - Data format translation (e.g., EBCDIC to ASCII)
   - Data encryption/decryption (e.g., TLS/SSL)
   - Data compression (reducing file sizes for transmission)
- Examples of Presentation Layer technologies:
   - SSL/TLS (Secure Sockets Layer/Transport Layer Security)
   - JPEG, GIF, MPEG (file formats for images/videos)

##### Application Layer (Web & Services)
- Purpose:
   - The Application Layer is the closest to the user. It provides network services directly to user applications (such as email, file transfer, web browsing).
- Key Responsibilities:
   - Providing network services to users (browser, email client)
   - Identifying communication partners (figuring out where to send)
   - Authentication and privacy services
- Examples of Application Layer technologies:
   - HTTP/HTTPS (web browsing)
   - SMTP (email sending)
   - FTP (file transfer)
   - DNS (domain name resolution)

##### How It All Works Together
- When you type example.com in your browser, a DNS request finds the website’s IP address.
- Your device sends packets of data via routers and switches across various networks.
- These packets use IP addresses to determine the destination.
- TCP/UDP protocols ensure proper delivery.
- Once the data reaches the server, it responds with web page data that is sent back following the same process.

#### How DNS Works?
- DNS works efficiently, translating user-friendly domain names into IP addresses, allowing seamless navigation on the internet.
- Below step by step working of DNS:
  - User Input: When a user enters a domain name in a browser, the system needs to find its IP address.
  - DNS Query: The user’s device sends a DNS query to the DNS resolver.
  - Resolver Request: The DNS resolver checks its cache for the IP address. If not found, it forwards the request to the root DNS server.
  - Root DNS Server: The root DNS server provides the address of the TLD (Top-Level Domain) server for the specific domain extension (e.g., .com).
  - TLD DNS Server: The TLD server directs the resolver to the authoritative DNS server for the actual domain.
  - Authoritative DNS Server: The authoritative DNS server knows the IP address for the domain and provides it to the resolver.
  - Response to User: The resolver stores the IP address in its cache and sends it to the user’s device.
  - Access Website: With the IP address, the user’s device can access the desired website.
  
### Glossary
#### Difference Between Stateless and Stateful Protocol
- Stateless Protocols are the type of network protocols in which the Client sends a request to the server and the server responds back according to the current state. It does not require the server to retain session information or status about each communicating partner for multiple requests. 
   - HTTP (Hypertext Transfer Protocol), UDP (User Datagram Protocol), and DNS (Domain Name System) are examples of Stateless Protocols. 
- Salient Features of Stateless Protocols
   - Stateless Protocol simplifies the design of the Server.
   - The stateless protocol requires fewer resources because the system does not need to keep track of the multiple link communications and the session details.
   - In Stateless Protocol each information packet travels on its own without reference to any other packet.
   - Each communication in Stateless Protocol is discrete and unrelated to those that precedes or follow.
- In Stateful Protocol If client send a request to the server then it expects some kind of response, if it does not get any response then it resend the request. FTP (File Transfer Protocol), TCP, and Telnet are the example of Stateful Protocol. 
- Salient Features of Stateful Protocol
   - Stateful Protocols provide better performance to the client by keeping track of the connection information.
   - Stateful Application require Backing storage.
   - Stateful request are always dependent on the server-side state.
   - TCP session follow stateful protocol because both systems maintain information about the session itself during its life.
