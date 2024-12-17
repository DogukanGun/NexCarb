# NexCarb - Real-Time Data Collection Network

## Project Aim

The aim of this project is to:
- Collect real-time data via a **distributed network**.
- Use this collected data for **oracles** and **AI models** in the near future.
- Replace **corrupted middlemen** with people who directly collect data and earn money for their contributions.

---

## Process

### Node Types

There are **two types of nodes** used to collect data:

1. **Sensor Nodes**
   - Sensor nodes work on vehicles.
   - A user orders a device and attaches it to their vehicle.
   - Once activated, the device starts collecting real-time data.
   - After reaching the maximum data limit, the device **sells the data** to the network.
   - Firms or individuals can buy this data from the network marketplace.

2. **Search Nodes**
   - Search nodes are lightweight and operate on **users' phones**.
   - A search node connects to the network and manages **multiple devices**.
   - When a data request comes to the network (e.g., from a firm or individual), the search node:
     - Randomly selects connected phones.
     - Responds to the selected phone's query with data instructions.
     - The phone fetches the required data in the background and sends it back to the network.

---

### Additional Nodes

- **Manager Node**:  
  A bootstrap node that stores and manages information about other nodes in the network.

- **AI Node** *(Upcoming)*:  
  An AI node will be introduced to train and run **AI models** using the collected data.

- **Main Oracle**:  
  The oracle node processes and utilizes the collected data inside the network.

---

## UI

### Sensor Node UI
- Sensor nodes have a **user interface (UI)** that users can access via a network connection.
- This page displays:
  - **Current energy consumption**
  - **Collected data**  
- **Future Goal**:  
  The device will support multiple sensors that can be plugged in and out, allowing collection of **different types of data** from the vehicle.

### Search Node CLI
- Search nodes are controlled via a **command-line interface (CLI)**.
- **User Perspective**:  
  Users interact with a **quiz application**:
  - Users can **create quizzes** and share them with friends.
  - Friends solve the quizzes while the app silently fetches data in the background.

---

## Earning Mechanism

The network enables data collectors and contributors to earn revenue as follows:

1. **Data Collection**:  
   Users earn money by selling the data they collect.

2. **Data Marketplace**:  
   Firms or individuals can purchase datasets directly from the network's marketplace.

3. **Revenue Distribution**:  
   Revenue generated from each dataset purchase is distributed to:
   - **Network operators**
   - **Data collectors**
   - **The platform**

4. **Payment Options**:  
   Datasets can be purchased using:
   - **USDC** (USD Coin)
   - **The network's native token**

---

## Summary

This project decentralizes real-time data collection and incentivizes individuals to contribute data using distributed nodes. By cutting out middlemen, the platform ensures transparency and rewards data collectors fairly.
