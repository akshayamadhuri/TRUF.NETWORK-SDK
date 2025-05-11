
### **Key Features**

1. **Decentralized Platform**:
   - The TRUF Network is a decentralized platform designed for interacting with economic data streams.
   - It supports the deployment and interaction with smart contracts through a `streamID`.

2. **Economic Data Streams**:
   - The platform specializes in **economic data streams**, making it a valuable resource for financial and economic analyses.

3. **Stream Types**:
   - **Primitive Streams**: Direct data sources like economic indexes or sentiment analysis.
   - **Composed Streams**: Aggregated data from multiple streams.

4. **Stream Identification**:
   - Streams are identified by a unique `streamID`, which corresponds to a deployed contract.
   - These `streamID`s allow users to locate and interact with specific streams.

5. **Data Access**:
   - Developers can read data from streams using **date ranges** or **timestamps**.
   - Recent updates (Contract Version 2) introduced support for timestamps, enabling more precise data handling.

6. **SDKs**:
   - **Types**: SDKs are available in Go and TypeScript/JavaScript.
   - **Purpose**: These SDKs facilitate integration with the decentralized platform and provide tools to publish, compose, and consume economic data streams.
   - **Interaction**: Users can interact with streams via their `streamID`. Examples include loading streams and reading data.

7. **Timestamp Support**:
   - **Availability**: Support for timestamps has been added in **contract version 2**.
   - **Implementation**: This involves replacing the `Date` type with the `number` type in relevant structures like `GetRecordInput` and `Record`.
   - **Usage (Reading)**: You can read data from a stream using timestamps (numbers) for the date range parameters in `GetRecordInput`.
   - **Usage (Deployment)**: New streams can be deployed with timestamp support by using functions like `deployPrimitiveStream` and setting the `contractVersion` parameter to 2.

8. **Ecosystem Roles**:
   - **Data Providers**: Publish and compose streams of data.
   - **Consumers**: Analyze and utilize the data streams.
   - **Node Operators**: Maintain network infrastructure and participate in consensus mechanisms.
