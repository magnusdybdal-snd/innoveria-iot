import { useState } from "react";
import Box from "@mui/material/Box";
import Path from "../templates/path.tsx";
import SensorsGenInfo from "../templates/sensorsGenInfo.tsx";
import SensorInfo from "../templates/sensorInfo.tsx";
import Menu from "../Menu.tsx";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Button from "@mui/material/Button";

export default function Sensors() {
  const [count, setCount] = useState(0);

  const sensorInfos = new Map<string, number>();
  sensorInfos.set("Total sensors", count);
  sensorInfos.set("Online sensors", 0);
  sensorInfos.set("Offline sensors", 0);
  sensorInfos.set("Last seen 24hr", 0);
  sensorInfos.set("Error last 24hr", 0);

  return (
    <div className="flex h-screen">
      <Menu />
      <Box
        sx={{
          backgroundColor: "primary.dark",
          color: "primary.main",
        }}
        className="flex-1 overflow-auto"
      >
        <Box
          sx={{
            paddingLeft: 2,
            paddingRight: 2,
            paddingTop: 2,
          }}
          className="flex-1 overflow-auto"
        >
          <Path />
          <div className="flex justify-between flex-wrap">
            <Typography variant="h4">Sensor devices</Typography>
            <Button
              variant="outlined"
              sx={{
                backgroundColor: "primary.main",
                color: "primary.dark",
                "&:hover": {
                  backgroundColor: "primary.main",
                },
                borderRadius: 2,
                textTransform: "none",
                fontSize: 20,
              }}
              onClick={() => setCount((count) => count + 1)}
            >
              Add device +
            </Button>
          </div>

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
          {/*True or false for sensor to be enabled or disabled*/}
          {Array.from({ length: count }).map((_, i) => (
            <SensorInfo key={i} number={i + 1} online={true} />
          ))}
        </Box>
      </Box>
    </div>
  );
}
