import * as DropdownMenu from '@radix-ui/react-dropdown-menu';
import { ReactNode } from "react"
import { Link } from 'react-router-dom';
import { CiCircleQuestion, CiCircleRemove } from "react-icons/ci"
import { useAuthAction } from '../../../store';

type DropDownAccountProps = {
    trigger: ReactNode
}

function DropDownAccount({ trigger }: DropDownAccountProps) {
    const actions = useAuthAction()

    return (
        <DropdownMenu.Root>
            <DropdownMenu.Trigger asChild>
                {trigger}
            </DropdownMenu.Trigger>

            <DropdownMenu.Portal>
                <DropdownMenu.Content className='p-4 rounded-xl bg-yellow-200'>
                    <DropdownMenu.Item className='flex items-center outline-none p-2 rounded-xl hover:bg-yellow-300'>
                        <span className='mr-2 my-1'>
                            <CiCircleQuestion size="24"/>
                        </span>
                        <Link to="/questions">Questions</Link>
                    </DropdownMenu.Item>
                    <DropdownMenu.Item className='flex items-center outline-none p-2 rounded-xl hover:bg-yellow-300'>
                        <span className='mr-2 my-1'>
                            <CiCircleRemove size="24"/>
                        </span>
                        <button onClick={() => {
                            actions.setId("")
                            actions.setUsername("")
                        }}>Logout</button>
                    </DropdownMenu.Item>
                    <DropdownMenu.Arrow className='fill-none' />
                </DropdownMenu.Content>
            </DropdownMenu.Portal>
        </DropdownMenu.Root>
    )
}

export default DropDownAccount