import Typography from "@mui/material/Typography";
import CircleIcon from "@mui/icons-material/Circle";

type InfoProps = {
  name: string;
  online?: boolean;
  euid: string;
  lastSeen: string;
};

export function GatewayInfo({ name, online, euid, lastSeen }: InfoProps) {
  return (
    <>
      <CircleIcon
        color={online ? "success" : "error"}
        sx={{ fontSize: 14, alignSelf: "center" }}
      />
      <Typography>{name}</Typography>
      <Typography>{euid}</Typography>
      <Typography>{lastSeen}</Typography>
    </>
  );
}
