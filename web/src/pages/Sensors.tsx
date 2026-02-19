import { useState } from "react";

import Button from "@mui/material/Button";

import { PageContent } from "@/components/pageContent";
import { PageDivider } from "@/components/pageDivider";
import { SensorInfo } from "@/components/sensorInfo";
import { SensorsGenInfo } from "@/components/sensorsGenInfo";
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
      <PageContent>
        <SubPageHeader title="Sensor devices" action={addButton} />
        <div className="flex justify-between flex-wrap">
          {Array.from(sensorInfos.entries()).map(([key, value]) => (
            <SensorsGenInfo key={key} title={key} count={value} />
          ))}
        </div>
        <PageDivider />
        {Array.from({ length: count }).map((_, i) => (
          <SensorInfo key={i} number={i + 1} online={true} />
        ))}
      </PageContent>
    </div>
  );
}
