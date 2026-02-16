import Box from "@mui/material/Box";
import Link from "@mui/material/Link";
import Typography from "@mui/material/Typography";
import CircleIcon from "@mui/icons-material/Circle";

type InfoProps = {
  number: number;
  online: boolean;
};

{
  /*Format for single sensor info*/
}
export default function SensorsGenInfo({ number, online }: InfoProps) {
  return (
    <Link href="#" underline="none">
      <Box
        sx={{
          backgroundColor: "secondary.light",
          color: "primary.main",
          margin: 1,
          padding: 5,
          borderRadius: 3,
          border: "3px grey",
          display: "flex",
          flexDirection: "column",
        }}
      >
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            gap: 1,
          }}
        >
          <CircleIcon
            color={online ? "success" : "error"}
            sx={{ fontSize: 14 }}
          />
          <Typography fontSize={30} fontWeight={300}>
            Sensor {number}
          </Typography>
        </Box>
      </Box>
    </Link>
  );
}
