import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

type InfoProps = {
  title: string;
  count: number;
};

export function SensorsGenInfo({ title, count }: InfoProps) {
  return (
    <Box
      sx={{
        backgroundColor: "secondary.light",
        color: "primary.main",
        marginTop: 2,
        borderRadius: 3,
        border: "3px grey",
        width: "15%",
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
  );
}
