import { useState } from "react";
import Box from "@mui/material/Box";
import Divider from "@mui/material/Divider";
import Button from "@mui/material/Button";
import { SensorsGenInfo } from "@/components/sensorsGenInfo";
import { SensorInfo } from "@/components/sensorInfo";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu.tsx";

const sensorInfos = new Map<string, number>();
sensorInfos.set("Total sensors", 0);
sensorInfos.set("Online sensors", 0);
sensorInfos.set("Offline sensors", 0);
sensorInfos.set("Last seen 24hr", 0);
sensorInfos.set("Error last 24hr", 0);

export default function Sensors() {
  const [count, setCount] = useState(0);

  const addButton = (
    <Button
      variant="outlined"
      sx={{
        backgroundColor: "primary.main",
        color: "primary.dark",
        "&:hover": { backgroundColor: "primary.main" },
        borderRadius: 2,
        textTransform: "none",
        fontSize: 20,
      }}
      onClick={() => setCount((count) => count + 1)}
    >
      Add device +
    </Button>
  );

  return (
    <div className="flex h-screen">
      <Menu />
      <Box
        sx={{ backgroundColor: "primary.dark", color: "primary.main" }}
        className="flex-1 overflow-auto"
      >
        <Box sx={{ paddingLeft: 2, paddingRight: 2 }}>
          <SubPageHeader title="Sensor devices" action={addButton} />
          <div className="flex justify-between flex-wrap">
            {Array.from(sensorInfos.entries()).map(([key, value]) => (
              <SensorsGenInfo key={key} title={key} count={value} />
            ))}
          </div>
          <Divider
            sx={{
              backgroundColor: "primary.main",
              marginTop: 2,
              marginBottom: 5,
            }}
          />
          {Array.from({ length: count }).map((_, i) => (
            <SensorInfo key={i} number={i + 1} online={true} />
          ))}
        </Box>
      </Box>
    </div>
  );
}
