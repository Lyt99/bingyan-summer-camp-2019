# Docker notes

## Basic Concepts

- **Image**
- **Container**
- **Repository**

## Docker Image

OS = Kernel + user_space

After the start of the OS Kernel, it will load `root` file system to support the user space. **Docker Image** is like a `root` file system. It contains programs, packages, resources, settings and other parameters (anonymous volumes, environmental variables, users) that are needed for a container. Images don't possess any dynamic data and will remain the same after construction.

#### Tiered storage

File systems like `root` can be huge. Docker uses tiered storage (Union FS). This image is not like ISO files. It consists of a number of layers of file system. The lower layer is the basis of the higher layer. Once it is finished, it will not change. Even if some files are deleted, they are just *tagged* to be deleted. So, each layer should only contain needed things.

This tiered storage frame makes reusing and customization easier. Developers can use established layers as foundation.

## Docker Container

**Image** + **Container** = **Class** + **Instance**

Container is actually a special process. It is different from other processes in the host. It has its own namespace, thus a container can have its `root` file system, processes space, user ID space. It's also separated from its host.

Like image, container also uses tiered storage tech. Containers build container storage layers used for read/write on image layers. It has the same life span as the container.

Container storage layer should remain stateless. I/O should use **Volume**s or bind to host's directories. This tech read/write directly to the host, which is more stable and faster.

**Volume**s' lifespan is independent from containers.

## Docker Registry

This is used to store/distribute Images. One **Docker Registry** can contain several **Repositories**, with each can have several **tags**, which represent Images. 