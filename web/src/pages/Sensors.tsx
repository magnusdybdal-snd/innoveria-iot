import Button from "@mui/material/Button";
import { SensorsGenInfo } from "@/components/sensorsGenInfo";
import { SensorInfo } from "@/components/sensorInfo";
import { SubPageHeader } from "@/components/subPageHeader";
import Menu from "@/Menu.tsx";
import { PageContent } from "@/components/pageContent";
import { PageDivider } from "@/components/pageDivider";
import { CategoryHeader } from "@/components/CategoryHeader";
import { DeviceRow } from "@/components/gatewayRow";
import { mockSensors } from "@/mocks/sensors.ts";

const sensorInfos = new Map<string, number>();
sensorInfos.set("Total sensors", 0);
sensorInfos.set("Online sensors", 0);
sensorInfos.set("Offline sensors", 0);
sensorInfos.set("Last seen 24hr", 0);
sensorInfos.set("Error last 24hr", 0);

const sensorDetails: string[] = ["Status", "Name", "DebEUI", "Machine"];

export default function Sensors() {
  //const [count, setCount] = useState(mockSensors.length);

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
      //onClick={() => setCount((count) => count + 1)}
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
        <CategoryHeader
          categories={sensorDetails}
          columns={sensorDetails.length + 1}
        >
          {/*Array.from({ length: count }).map((_, i) => (
          <SensorInfo key={i} number={i + 1} online={true} />
        ))*/}
          {mockSensors.map((sensor) => (
            <DeviceRow key={sensor.id}>
              <SensorInfo
                name={sensor.name}
                status={sensor.status}
                euid={sensor.euid}
                machine={sensor.machine}
              />
            </DeviceRow>
          ))}
        </CategoryHeader>
      </PageContent>
    </div>
  );
}
