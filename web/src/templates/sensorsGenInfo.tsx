import Box from "@mui/material/Box";
import Link from "@mui/material/Link";
import Typography from "@mui/material/Typography";

type InfoProps = {
  title: string;
  count: number;
};

{
  /*Format for general sensors info*/
}
export default function SensorsGenInfo({ title, count }: InfoProps) {
  return (
    <Link href="#" underline="none">
      <Box
        sx={{
          backgroundColor: "secondary.light",
          color: "primary.main",
          margin: 1,
          borderRadius: 3,
          border: "3px grey",
          width: 250,
          height: 100,
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          textAlign: "center",
        }}
      >
        <div className="w-64 p-[10px]">
          <Typography>{title}</Typography>
          <Typography>{count}</Typography>
        </div>
      </Box>
    </Link>
  );
}
