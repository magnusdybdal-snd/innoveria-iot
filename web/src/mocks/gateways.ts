/**
 *
 */
export type Gateway = {
  id: string;
  name: string;
  status: number;
  euid: string;
  lastSeen: string;
};

export const mockGateways: Gateway[] = [
  {
    id: "1",
    name: "Gateway 1",
    status: 0,
    euid: "A1B2C3D4E5F6",
    lastSeen: "2 min ago",
  },
  {
    id: "2",
    name: "Gateway 2",
    status: 2,
    euid: "B2C3D4E5F6A1",
    lastSeen: "1 hour ago",
  },
  {
    id: "3",
    name: "Gateway 3",
    status: 1,
    euid: "B2D3D5E5A6A1",
    lastSeen: "4 days ago",
  },
];
