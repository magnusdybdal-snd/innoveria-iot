import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

type InfoProps = {
  title: string;
  count: number;
};

{
  /*Format for general sensors info*/
}
/**
 * Summary card displaying a sensor statistic label and its numeric count.
 * @param root0 - Component props
 * @param root0.title - Label for the statistic (e.g. "Online sensors")
 * @param root0.count - Numeric value of the statistic
 * @returns The rendered summary card
 */
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
