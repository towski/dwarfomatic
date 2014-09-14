#define LINUX_BUILD
#include <iostream>
#include <stdint.h>
#include "DataDefs.h"
#include "Export.h"
#include "RemoteClient.h"
#include "RemoteTools.h"
#include "BasicApi.pb.h"
#include "foo.hpp"
//#include "library/include/df/item.h"
dfproto::ListUnitsIn *in;
dfproto::ListUnitsOut *out;

void cxxFoo::Init(void) {
    DFHack::color_ostream_wrapper * df_network_out = new DFHack::color_ostream_wrapper(std::cout);
    DFHack::RemoteClient * network_client = new DFHack::RemoteClient(df_network_out);
    std::cout << network_client->connect() << "\n";
    DFHack::RemoteFunction<dfproto::ListUnitsIn, dfproto::ListUnitsOut> my_call;
    my_call.bind(network_client, "ListUnits");
    in = new dfproto::ListUnitsIn();
    out = new dfproto::ListUnitsOut();
    in->set_scan_all(true);
    //my_call(network_client->default_output(), dfproto::ListUnitsIn::default_instance, dfproto::ListUnitsOut::default_instance);
    my_call(in, out);
    std::cout << in->id_list_size() << std::endl;
    std::cout << out->value_size() << std::endl;
    //network_client->run_command((*df_network_out), "ls", std::vector<std::string>()); 
    //std::cout << network_client->suspend_game() << "\n";
    network_client->disconnect();

	std::cout<<this->a<<std::endl;
}

const char* cxxFoo::Bar(int i) {
    return out->value(i).name().last_name().c_str();
}

