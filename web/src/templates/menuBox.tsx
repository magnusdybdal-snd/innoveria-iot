import Box from "@mui/material/Box"
import Link from "@mui/material/Link"
import { Link as RouterLink } from 'react-router'

type SubPage = {
    name: string
    path: string
}

type MenuBoxProps = {
    title: string
    pages: SubPage[]
    add: boolean
};

{/*Box for menu components*/}
export default function MenuBox({ title, pages, add }: MenuBoxProps) {
    return (
        <Box
            sx={{
                backgroundColor: 'secondary.light',
                color: 'primary.main',
                margin: 1,
                borderRadius: 3,
                width: 300
            }}
        >
            <div className="w-64 p-[10px]">
                {/*Main page*/}
                <h2 className="font-normal text-3xl"  style={{ margin: 0 }}>
                    {title}
                </h2>

                <nav className="flex flex-col mt-2 space-y-1">
                    {pages.map((page) => (
                        <Link
                            key={page.path}
                            component={RouterLink}
                            to={page.path}
                            underline="none"
                            sx={{color: 'primary.main'}}
                        >
                            | {page.name}
                        </Link>
                    ))}

                    {add && (
                        <Link underline="none" sx={{color: 'primary.main'}}>
                            + Add
                        </Link>
                    )}
                </nav>
            </div>
        </Box>
    );
}