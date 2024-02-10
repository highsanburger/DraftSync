# IOT ECE352E Project

_Members_ : Syed Khalid 21BEE1288, Arya Nag 21BEE1336

## Aim :-

To develop a Peer-to-Peer (P2P) file synchronization web app using the BYOD (Bring Your Own Device) Cloud model.

[GitHub Link](https://github.com/highsanburger/DraftSync)

---

## User Story

As a user I will be able to :-

- open the webapp on any device (PC, Laptop, Android, iOS )
- connect as a node in the local network and discover other connected nodes over the local network
- upload files which can be downloaded by other nodes
- download files which are uploaded by other nodes
- synchronize two versions of a file shared by two nodes

## Screen Flow

Home Page -> Connect to Network (Broadcasts your Node) -> Upload Files / Downlaod Files / Select Node -> Sync Files -> Settings / Disconnect

---

## Technologies Used :-

- **Backend:** Golang
- **Frontend:** HTMX
- **Styling:** Tailwind CSS

### Why these technologies?

- Golang for efficient backend development with minimum boilerplate from external frameworks.
- HTMX for dynamic and interactive frontend with straightforward development cycle without the bloat of 3rd party libraries.
- Tailwind CSS for a clean and responsive UI.

---

## Progress :-

- Set up all the basics and have a bare minimum webapp.

- Conducted research on relevant Go frameworks:
  - Explored options for P2P communication (go-libp2p or peerdiscovery).
  - Investigated HTTP/TCP standard libraries for developing the webapp (net/http).
  - Familiarising the building process of the frontend using plain HTML. (HTMX & html/template )
